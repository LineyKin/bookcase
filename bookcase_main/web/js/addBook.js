$( document ).ready(function(){
    console.log("bookcase_jwt: "+$.cookie("bookcase_jwt"))
    console.log($.cookie("bookcase_jwt")==null)
    if($.cookie("bookcase_jwt") == null) {
        window.location.replace("/auth")
    }
    getAuthorHint()
    getPublishingHouseHint()
    buildPaginator()
    getBookList(1, "author", "asc", true)
});


function createOptionForDatalist(id, name) {
    return '<option name="'+name+'" data-id="'+id+'">'+name+'</option>';
}

// кнопка, добавляющая ещё одно поле для автора, если у книги несколько авторов
$("#addAuthorBtn").on("click", function() {
    let authorRow = $(this).parent().parent()
    authorRow.after(`<tr>
        <td></td>
        <td>
            <input type="text" list="authors" class="author">
            <button class="deleteAuthorInput" type="button" class="btn btn-primary">-</button>
        </td>
    </tr>`)

    $(".deleteAuthorInput").on("click", function() {
        $(this).parent().parent().remove()
    });
});

// добавления поля с произведением, если одна физическая кника содержит более одного произведения
// например "Час быка" и "Туманность Андромеды" - два романа Ефремова в одной физической книге
$("#addWorkBtn").on("click", function() {
    let workRow = $(this).parent().parent()
    workRow.after(`<tr>
        <td></td>
        <td>
            <input type="text" class="literaryWork">
            <button class="deleteWorkInput" type="button" class="btn btn-primary">-</button>
        </td>
    </tr>`)

    $(".deleteWorkInput").on("click", function() {
        $(this).parent().parent().remove()
    });
});

// сохранение (физической) книги
$("#saveBook").on("click", function(){

    // получение списка id авторов
    let authorIdList = [];
    $(".author").each(function(){
        let authorId = $('#authors [name = "'+$(this).val()+'"]').data("id");
        if(typeof(authorId) == "undefined") {
            authorId = 0;
        }
        authorIdList.push(authorId)
    });

    // получение списка литературных произведений.
    // элемент списка - ассоциативный массив id-name.
    // если произведение добавляется впервые, то id=0
    //
    // id нужно для 
    // - связки одного произведения с разными физическими носителями
    // например, "Чук и Гек" изданый "Детиздатом" в 1953 и изданый АСТ в 2024.
    // - для различения книг разных авторов с одним названием, например "Немезида" А.Кристи и А.Азимова
    // 
    let literaryWorkList = [];
    $(".literaryWork").each(function(){
        let literaryWorkName = $(this).val()
        let literaryWorkId = $('#literaryWorks [name = "'+literaryWorkName+'"]').data("id")
        if(typeof(literaryWorkId) == "undefined") {
            literaryWorkId = 0;
        }

        let literaryWork = {
            id: literaryWorkId, 
            name: literaryWorkName
        }
        
        literaryWorkList.push(literaryWork)
    });

    // получение издательства
    let publishingHouseName = $("#publishingHouse").val()
    let publishingHouseId = $('#publishingHouses [name = "'+publishingHouseName+'"]').data("id")
    if(typeof(publishingHouseId) == "undefined") {
        publishingHouseId = 0
    }
    let publishingHouse = {
        id : publishingHouseId,
        name: publishingHouseName
    }

    // получение года
    let publishingYear = $("#publishingYear").val()

    // данные для отправки в post-запрос
    let bookData = {
        author: authorIdList,
        name: literaryWorkList,
        publishingHouse: publishingHouse,
        publishingYear: publishingYear,
        jwt: $.cookie("bookcase_jwt")
    }
    $.ajax({
        type: "POST",
        url: "api/book/add",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(bookData),
        success: function (response) {
            console.log(response)
            // очистка формы
            $(".author").each(function(){
                $(this).val("")
            });

            $(".literaryWork").each(function(){
                $(this).val("")
            });

            $("#publishingHouse").val("")
            $("#publishingYear").val("")
        },
        error: function (errorResponse) {
            console.log(errorResponse)
        }
    })
});

// добавление автора
$("#saveAuthor").on("click", function() {
    let authorData = {
        name: $("#authorName").val(),
        fatherName: $("#authorFatherName").val(),
        lastName: $("#authorLastName").val()
    };

    $.ajax({
        type: "POST",
        url: "api/author/add",
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(authorData),
        success: function (response) {

            console.log(response)
        
            // очистка формы
            $("#authorName").val(""),
            $("#authorFatherName").val(""),
            $("#authorLastName").val("")

            // сообщение, что автор добавлен
            let message = "Автор " + authorData.name + " " + authorData.lastName + " добавлен"
            console.log(message)
            $("#addAuthSuccessMessage").show()
            $("#addAuthSuccessMessage").html(message)
            $("#addAuthSuccessMessage").hide(20000)

            // обновление datalist
            let authorId = response.author_id
            let author = authorData.name +" "+ authorData.lastName
            optionTag = '<option name="'+author+'" data-id="'+authorId+'">'+author+'</option>';
            $("#authors").append(optionTag)

        },
        error: function (errorResponse) {
            let status = errorResponse.status + " " + errorResponse.statusText
            let errorText = errorResponse.responseJSON.error
            let message = "Ошибка добавления в БД. Статус: " + status + ". Ошибка: " + errorText
            console.log(message)
            $("#addAuthErrorMessage").show()
            $("#addAuthErrorMessage").html(message)
            $("#addAuthErrorMessage").hide(20000)
        }
    });
});