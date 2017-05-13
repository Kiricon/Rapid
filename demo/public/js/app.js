(function(){
    let button = document.getElementById('test');

    button.addEventListener('click', () => {
        let xhr = new XMLHttpRequest();
        let url = "/post";
        xhr.open("POST", url, true);
        xhr.setRequestHeader("Content-type", "application/json");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                var json = JSON.parse(xhr.responseText);
                console.log(json.email + ", " + json.password);
            }
        };
        let data = JSON.stringify({"email": "hey@mail.com", "password": "101010"});
        xhr.send(data);
    });
})()