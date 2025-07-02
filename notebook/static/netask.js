const ipConfig = "http://47.95.200.35:8089";

export async function addReminderToBack(num, text) {

    const response = await fetch(ipConfig + "/add", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            ID: num,
            Data: text,
            Done: false
        })
    });

    if(!response.ok){
        throw new Error("add failed, Net word error");
    }

    const data = await response.json();

    if(data.message == "success"){
        return true;
    }
    else {
        throw new Error(data.message)
    }
}

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

    if(!response.ok){
        throw new Error("get failed, Net word error");
    }

    const data = await response.json();

    if(data.message == "success" || data.message == "this user has no more note"){
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

    console.log(num);

    if(!response.ok){        
        throw new Error("updata failed, Net word error");
    }

    const data = await response.json();

    if(data.message == "updata success"){
        return true;
    }
    else {
        throw new Error(data.message);
    }
}