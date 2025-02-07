$("#registerBtn").on("click", function(){
    $.ajax({
        type: "POST",
        url: "register",
        success: function (response) {
            console.log(response)
            window.location.replace("/")
        },
        error: function (errorResponse) {
            console.log(errorResponse)
        }
    });
})