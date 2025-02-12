$("#loginBtn").on("click", function(){
    let loginData = {
        login: $("#login").val(),
        password: $("#password").val(),
    };

    $.ajax({
        type: "POST",
        url: "login",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(loginData),
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

$("#logoutBtn").on("click", function() {
    $.removeCookie('bookcase_jwt')
    window.location.replace("/auth")
});