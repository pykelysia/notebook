
import { addReminderToBack, updata, getReminderFromBack } from "./netask.js";

let num = 1;

getReminder();

async function getReminder() {
    while(true){
        const data = await getReminderFromBack(num);

        if(data.message == "this user has no more note")
            break;
        console.log(data.data);
        const list = document.querySelector('.reminders-list');
        const newItem = document.createElement('div');
        newItem.className = 'reminder-item';
        newItem.innerHTML = `
            <div class="checkbox"></div>
            <div class="reminder-text" id ="${data.ID}">${data.data}</div>`;
        newItem.querySelector('.checkbox').addEventListener('click', beClick);
        list.prepend(newItem);
        num+=1;
    }
}

// 写入Add
function addReminder() {

    const input = document.getElementById('newReminder');
    const text = input.value.trim();
    
    if (text) {
        console.log(num);
        addReminderToBack(num, text).then((success) => {
            if(!success){
                return;
            }
            const list = document.querySelector('.reminders-list');
            const newItem = document.createElement('div');
            newItem.className = 'reminder-item';
            newItem.innerHTML = `
                <div class="checkbox"></div>
                <div class="reminder-text" id ="${num}">${text}</div>`;

            newItem.querySelectorAll('.checkbox').forEach(checkbox => {
                checkbox.addEventListener('click', beClick);
            })

            list.prepend(newItem);
            input.value = '';
            num = num + 1;
            })
    }
}

// 添加回车键支持
document.getElementById('newReminder').addEventListener('keypress', function(e) {
    if (e.key === 'Enter') {
        addReminder();
    }
});
document.getElementById('addReminder').addEventListener('click',addReminder);

// 初始化示例事项的点击事件    
// 涉及 Update 
document.querySelectorAll('.checkbox').forEach(checkbox => {
    checkbox.addEventListener('click', beClick);
});

function beClick(){
    if(this.classList.contains("checked")){
        updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, true).then((success) => {
            this.classList.toggle('checked');
            this.nextElementSibling.classList.toggle('completed');
        })
    }
    else{
        updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, false).then((success) => {
            this.classList.toggle('checked');
            this.nextElementSibling.classList.toggle('completed');
        })
    }
}