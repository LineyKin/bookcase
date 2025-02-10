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
        success: function (response) {
            console.log(response.jwt)
            alert(response.jwt)
            //window.location.replace("/")
        },
        error: function (errorResponse) {
            console.log(errorResponse)
        }
    });
})