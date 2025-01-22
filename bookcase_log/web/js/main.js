$( document ).ready(function(){
    console.log("main.js here")
    buildPaginator()
});

// количество строк на странице
const rowsLimit = 10

// количество кнопок пагинатора
let paginatorItemCount = 0

function buildPaginator() {
    let bookCount = getLogCount()
    let num = bookCount/rowsLimit
    paginatorItemCount = num%10 == 0? num:  Math.trunc(num) + 1
    for (let i=1; i<= paginatorItemCount; i++) {
        let newItem = '<li class="page-item"><a class="page-link" href="#">' + i + '</a></li>'
        $("#bookListPagination").append(newItem)
    }

    $("#bookListPagination .page-item").on("click", function(){
        let paginatorNumber = $(this).find("a").html()

        let sortedBy, sortType
        $("#bookListTable th").each(function(){
            if($(this).attr("isSorted") != undefined) {
                sortedBy = $(this).attr("name")
                sortType = $(this).attr("isSorted")
            }
        })

        getBookList(paginatorNumber, sortedBy, sortType)
    })
}

function getLogCount() {
    let count = 0;
    $.ajax({
        type: "GET",
        async: false,
        url: "api/log/count",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            count = response.count

        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка выгрузки количества книг. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });

    return count
}