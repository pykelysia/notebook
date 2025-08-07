const ipConfig = "http://localhost:8089";

const url = new URL(window.location.href);
const searchParams = new URLSearchParams(url.search);
const uid = parseInt(searchParams.get('uid'), 10);

check();

function check(){
    checkUID(uid).then( (e) => {
        if(e == false){
            window.location.href = ipConfig + '/user';
        }
    })
}

async function checkUID(uid) {

    const response = await fetch(ipConfig + "/user/check", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            uid: uid,
        })
    });

    if (!response.ok) {
        throw new Error("Net word error");
    }

    const data = await response.json();

    if(data.message == "right") {
        return true;
    }
    return false;
}