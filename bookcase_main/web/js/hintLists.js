// выгрузка списка авторов в datalist
function getAuthorHint() {
    $.ajax({
        type: "GET",
        url: "api/author/hint",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            let data = response.author_list
            let authorCount = data.length
            for (let i=0; i< authorCount; i++) {
                
                let authorId = data[i].id
                let author = data[i].name +" "+ data[i].lastName
                optionTag = createOptionForDatalist(authorId, author);
                $("#authors").append(optionTag)
            }
        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка выгрузки списка-подсказки авторов. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });
}
function getPublishingHouseHint() {
     // выгрузка списка издательств в datalist
     $.ajax({
        type: "GET",
        url: "api/publishingHouse/list",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            let data = response.ph_list
            let phCount = data.length
            for (let i=0; i< phCount; i++) {
                optionTag = createOptionForDatalist(data[i].id, data[i].name);
                $("#publishingHouses").append(optionTag)
            }
        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка выгрузки списка-подсказки издательств. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            alert(message)
        }
    });
}
