(function () {
    let current: HTMLOrSVGScriptElement = document.currentScript,
        interval: string = current.getAttribute("interval") || "2000";

    const loop = () => {
        let xhr: XMLHttpRequest = new XMLHttpRequest();

        xhr.open('GET', 'http://localhost:8080/counter', true)

        let token = localStorage.getItem("token")
        if (token != null) xhr.setRequestHeader("Authorization", "Bearer " + token)
        xhr.onload = () => {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    let res = JSON.parse(xhr.responseText)
                    if (res.success === true) {
                        let data = res.data;
                        document.getElementById("online_user").innerHTML = data.online_user;
                        document.getElementById("online_total").innerHTML = formatTime(data.online_total);
                        document.getElementById("online_me").innerHTML = formatTime(data.online_me);

                        // set token
                        let setToken = xhr.getResponseHeader("Set-Token")
                        if (token == null && setToken != null) {
                            localStorage.setItem("token", setToken)
                        }
                    } else {
                        console.error(res.message)
                    }
                    setTimeout(loop, parseInt(interval))
                }
            }
        }
        xhr.send()
    }

    // secound formater
    const formatTime = (time: number) => {
        let day = Math.floor(time / (60 * 60 * 24));
        let hour = Math.floor((time % (60 * 60 * 24)) / (60 * 60));
        let minute = Math.floor((time % (60 * 60)) / 60);
        let second = Math.floor(time % 60);
        return `${day}d ${hour}h ${minute}m ${second}s`;
    }

    loop()
})()