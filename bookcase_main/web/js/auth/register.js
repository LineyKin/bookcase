$("#registerBtn").on("click", function(){
    let registerData = {
        login: $("#login").val(),
        password: $("#password").val(),
    };

    $.ajax({
        type: "POST",
        url: "register",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(registerData),
        success: function () {
            window.location.replace("/")
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