$( document ).ready(function(){
    buildPaginator()
    getLogList(1)
});

// количество строк на странице
const rowsLimit = 10

// количество кнопок пагинатора
let paginatorItemCount = 0

function getLogList(paginatorNumber) {

    let offset = rowsLimit * (paginatorNumber - 1)

    // Выгрузка списка логов
    $.ajax({
        type: "GET",
        data: {
            limit: rowsLimit,
            offset: offset,

        },
        url: "api/log/list",
        success: function (response) {
            buildLogTable(response.log_list)
        },
        error: function (errorResponse) {
            console.log("error")
            console.log(errorResponse)
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseText
            let message = "Ошибка выгрузки списка логов. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });
}

function buildLogTable(bookListArray) {
    $("#logTable .logRow").remove()
    len = bookListArray.length
    for (let i=0; i < len; i++) {
        let newRow = buildLogRow(bookListArray[i])
        $("#logTable").append(newRow)
    }
}

function dateTimeFormat(ts) {
    ts = ts.replace("T", ' ')
    ts = ts.replace("Z", '')

    return ts.split(".")[0]
}

function buildLogRow(obj) {
    return `<tr class="logRow">
                <td>`+dateTimeFormat(obj.producer_ts)+`</td>
                <td>`+obj.message+`</td>
            </tr>`
}


function buildPaginator() {
    let logCount = getLogCount()
    let num = logCount/rowsLimit
    paginatorItemCount = num%10 == 0? num:  Math.trunc(num) + 1
    for (let i=1; i<= paginatorItemCount; i++) {
        let newItem = '<li class="page-item"><a class="page-link" href="#">' + i + '</a></li>'
        $("#logListPagination").append(newItem)
    }

    $("#logListPagination .page-item").on("click", function(){
        let paginatorNumber = $(this).find("a").html()

        getLogList(paginatorNumber)
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