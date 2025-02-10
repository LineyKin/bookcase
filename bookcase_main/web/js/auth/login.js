$("#loginBtn").on("click", function(){
    let loginData = {
        login: $("#login").val(),
        password: $("#password").val(),
    };

    console.log(loginData)
    $.ajax({
        type: "POST",
        url: "login",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(loginData),
        success: function (response) {
            console.log(response)
            //window.location.replace("/")
        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка идентификации пользователя: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(errorText)
        }
    });
})