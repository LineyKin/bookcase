$("#loginBtn").on("click", function(){
    $.ajax({
        type: "POST",
        url: "login",
        success: function (response) {
            console.log(response)
            //window.location.replace("/")
        },
        error: function (errorResponse) {
            console.log(errorResponse)
        }
    });
})