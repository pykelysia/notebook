
import { addReminderToBack, updata } from "./netask.js";

let num = 0;

// getReminder();

// function getReminder() {
//     while(true){
//         let opt = false;
//         getReminderFromBack(num).then((data) => {
//             if(data.message == "this user has no more note"){
//                 opt = true;
//                 return;
//             }
//             const list = document.querySelector('.reminders-list');
//             const newItem = document.createElement('div');
//             newItem.className = 'reminder-item';
//             newItem.innerHTML = `
//                 <div class="checkbox"></div>
//                 <div class="reminder-text" id ="${data.UID}">${data.Data}</div>`;
//                 newItem.querySelector('.checkbox').addEventListener('click', function() {
//                 this.classList.toggle('checked');
//                 this.nextElementSibling.classList.toggle('completed');
//             });
//             list.prepend(newItem);
//             input.value = '';
//         })
//         if(opt)
//             break;
//         num+=1;
//     }
// }

// 写入Add
function addReminder() {

    const input = document.getElementById('newReminder');
    const text = input.value.trim();
    
    if (text) {
        num+=1;
        addReminderToBack(num, text).then((success) => {
            if(!success){
                return;
            }
        })
        const list = document.querySelector('.reminders-list');
        const newItem = document.createElement('div');
        newItem.className = 'reminder-item';
        newItem.innerHTML = `
            <div class="checkbox"></div>
            <div class="reminder-text" id ="${num}">${text}</div>`;
            newItem.querySelector('.checkbox').addEventListener('click', function() {
                this.classList.toggle('checked');
                this.nextElementSibling.classList.toggle('completed');
                if(this.classList.contains("checked")){
                    updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, true).then((success) => {
                        console.log(success);
                    })
                }
                else{
                    updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, false).then((success) => {
                        console.log(success);
                    })
                }
        });
        list.prepend(newItem);
        input.value = '';
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
    checkbox.addEventListener('click', function() {
        this.classList.toggle('checked');
        this.nextElementSibling.classList.toggle('completed');
        if(this.classList.contains("checked")){
            updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, true).then((success) => {
                console.log(success);
            })
        }
        else{
            updata(parseInt(this.nextElementSibling.id, 10), this.nextElementSibling.innerHTML, false).then((success) => {
                console.log(success);
            })
        }
    });
});