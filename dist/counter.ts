!function () {
    let interval: number = 1000;

    const loop = () => {
        let xhr: XMLHttpRequest = new XMLHttpRequest()
        xhr.open('GET', 'http://localhost:8080/counter', true)
        xhr.onload = () => {
            if (xhr.status === 200) {
                let counter = JSON.parse(xhr.responseText)
                console.log(counter)
                setTimeout(loop, interval)
            }
        }
        xhr.send()
    }

    const formatTime = (time: number) => {
        let h = time / 3600
        let m = (time - h * 3600) / 60
        let s = time - h * 3600 - m * 60
        return h + 'h ' + m + 'm ' + s + "s"
    }

    loop()
}