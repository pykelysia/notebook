import "./netask.js"

import { getReminder, addReminderToBack, updata } from "./netask.js";

getReminder();

// 写入Add
function addReminder() {

    const input = document.getElementById('newReminder');
    const text = input.value.trim();
    
    if (text) {
        addReminderToBack(text).then((data) => {
            if(data.message != 'success'){
                return;
            }
            const list = document.querySelector('.reminders-list');
            const newItem = document.createElement('div');
            newItem.className = 'reminder-item';
            newItem.innerHTML = `
                <div class="checkbox"></div>
                <div class="reminder-text" id ="${data.num}">${text}</div>`;

            newItem.querySelectorAll('.checkbox').forEach(checkbox => {
                checkbox.addEventListener('click', beClick);
            })

            list.prepend(newItem);
            input.value = '';
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

export function beClick(){
    if(this.classList.contains("checked")){
        updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, true).then(() => {
            this.classList.toggle('checked');
            this.nextElementSibling.classList.toggle('completed');
        })
    }
    else{
        updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, false).then(() => {
            this.classList.toggle('checked');
            this.nextElementSibling.classList.toggle('completed');
        })
    }
}