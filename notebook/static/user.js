
// 切换登录/注册表单
const loginTab = document.getElementById('login-tab');
const signupTab = document.getElementById('signup-tab');
const loginForm = document.querySelector('.login-form');
const signupForm = document.querySelector('.signup-form');
const switchBtn = document.getElementById('switch-btn');
const switchText = document.getElementById('switch-text');

loginTab.addEventListener('click', () => {
    loginTab.classList.remove('active');
    signupTab.classList.remove('active');
    loginForm.classList.remove('active');
    signupForm.classList.remove('active');
    switchText.textContent = '没有账户?';
    switchBtn.textContent = '立即注册';
});

signupTab.addEventListener('click', () => {
    signupTab.classList.add('active');
    loginTab.classList.add('active');
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

// 表单提交处理（简化示例）
document.querySelector('.btn-primary').addEventListener('click', (e) => {
    e.preventDefault();
    if (loginTab.classList.contains('active')) {
        alert('登录请求已发送！');
    } else {
        alert('注册请求已发送！');
    }
});