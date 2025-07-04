const ipConfig = "http://localhost:8089";

import {beClick} from "./main.js"

export const url = new URL(window.location.href);
export const searchParams = new URLSearchParams(url.search);
export const uid = parseInt(searchParams.get('uid'), 10);

//添加date
export async function addReminderToBack(text) {

    console.log(uid);

    const response = await fetch(ipConfig + "/add", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            uid: uid,
            password: 0,
            code: ['\0'],
            data: text,
            done: false,
        })
    });

    if (!response.ok) {
        throw new Error("add failed, Net word error");
    }

    const data = await response.json();

    if (data.message == "success") {
        return data;
    }
    else {
        throw new Error(data.message)
    }
}

export async function getReminder() {
    //获取uid下的所有date的code
    getCode(uid).then( code => {
        //按code从back end提取dates
        for(let i in code){
            const data = getReminderFromBack(parseInt(code[i], 10)); //code : string
            getReminderFromBack(parseInt(code[i], 10)).then( data => {
                const list = document.querySelector('.reminders-list');
                const newItem = document.createElement('div');
                newItem.className = 'reminder-item';
                newItem.innerHTML = `
                    <div class="checkbox"></div>
                    <div class="reminder-text" id ="${data.id+''}">${data.data}</div>`;
                newItem.querySelector('.checkbox').addEventListener('click', beClick);
                list.prepend(newItem);
                const child = newItem.querySelector('.checkbox');
                if(data.done){
                    child.classList.toggle('checked');
                    child.nextElementSibling.classList.toggle('completed');
                }
            })
        }
    })
}

//获取code

async function getCode(uid) {

    const response = await fetch(ipConfig + "/user/get", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            UID: uid,
        })
    });

    if (!response.ok) {
        throw new Error("get failed, Net word error");
    }

    const data = await response.json();

    if (data.message == "success") {
        return data.code;
    }
    else {
        throw new Error(data.message)
    }
}

//获取date的请求函数
export async function getReminderFromBack(num) {

    const response = await fetch(ipConfig + "/get", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            ID: num,
        })
    });

    if (!response.ok) {
        throw new Error("get failed, Net word error");
    }

    const data = await response.json();

    if (data.message == "success" || data.message == "this user has no more note") {
        return data;
    }
    else {
        throw new Error(data.message)
    }
}

export async function updata(num, text, done) {


    const response = await fetch(ipConfig + "/updata", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            ID: num,
            Data: text,
            Done: done
        })
    });

    if (!response.ok) {
        throw new Error("updata failed, Net word error");
    }

    const data = await response.json();

    if (data.message == "updata success") {
        return true;
    }
    else {
        throw new Error(data.message);
    }
}