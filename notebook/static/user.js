const ipConfig = "http://localhost:8089";

// 切换登录/注册表单
const loginTab = document.getElementById('login-tab');
const signupTab = document.getElementById('signup-tab');
const loginForm = document.querySelector('.login-form');
const signupForm = document.querySelector('.signup-form');
const switchBtn = document.getElementById('switch-btn');
const switchText = document.getElementById('switch-text');

loginTab.addEventListener('click', () => {
    signupTab.classList.remove('active');
    loginTab.classList.add('active');
    loginForm.classList.remove('active');
    signupForm.classList.remove('active');
    switchText.textContent = '没有账户?';
    switchBtn.textContent = '立即注册';
});

signupTab.addEventListener('click', () => {
    loginTab.classList.remove('active');
    signupTab.classList.add('active');
    signupForm.classList.add('active');
    loginForm.classList.add('active');
    switchText.textContent = '已有账户?';
    switchBtn.textContent = '立即登录';
});

switchBtn.addEventListener('click', (e) => {
    e.preventDefault();
    if (loginTab.classList.contains('active')) {
        signupTab.click();
    } else {
        loginTab.click();
    }
});

//表单提交
document.getElementById('signin-action').addEventListener('click',e => {
    e.preventDefault();
    const uidInput = document.getElementById('uid');
    const uid = uidInput.value.trim();
    const passwordInput = document.getElementById('password');
    const password = passwordInput.value.trim();

    if(!uid){
        alert('uid 不能为空');
        return;
    }
    if(!password){
        alert('密码不能为空');
        return;
    }

    userSighin(parseInt(uid, 10), parseInt(password, 10)).then( e => {
        if(e)
            window.location.href = ipConfig + '/?uid=' + uid;
    })
})

document.getElementById('signup-action').addEventListener('click', e => {
    e.preventDefault();
    alert("signup");
    const uidInput = document.getElementById('signup-uid');
    const uid = uidInput.value.trim();
    const passwordInput = document.getElementById('signup-password');
    const password = passwordInput.value.trim();

    if(!uid){
        alert('uid 不能为空');
        return;
    }
    if(!password){
        alert('密码不能为空');
        return;
    }
    userSignup(parseInt(uid, 10), parseInt(password, 10)).then(e => {
        if(e)
            console.log("sign up success");
    })
})

//网络请求
async function userSignup(uid, password) {
    const response = await fetch(ipConfig + "/user/signup", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            uid: uid,
            password: password
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
async function userSighin(uid, password) {
    const response = await fetch(ipConfig + "/user/signin", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            uid: uid,
            password: password
        })
    });

    if(!response.ok){
        throw new Error("sign in failed, Net word error");
    }

    const data = await response.json();

    if(data.message == "success"){
        return true;
    }
    else {
        throw new Error(data.message)
    }
}