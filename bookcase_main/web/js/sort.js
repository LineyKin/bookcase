let asc = "asc"
let desc = "desc"

$("#bookListTable th").on("click", function(){
    if ($(this).attr("isSorted") == undefined) {
        // удаляем признаки сортировки на старом столбце
        $("#bookListTable th").each(function(){
            if($(this).attr("isSorted") != undefined) {
                $(this).find("#sortIcon").remove()
                $(this).removeAttr("isSorted")
            }
        });

        $(this).html($(this).html()+sortIconUp()) 
        $(this).attr("isSorted", asc)
    } else {
        $(this).find("#sortIcon").remove()
        if ($(this).attr("isSorted") == asc) {
            $(this).html($(this).html()+sortIconDown()) 
            $(this).attr("isSorted", desc)
        } else {
            $(this).html($(this).html()+sortIconUp()) 
            $(this).attr("isSorted", asc)
        }
    }

    isTotal = 0
    if ($("#bookListTable").attr("class") == "table table-striped total") {
        isTotal = 1
    }

    getBookList(1, $(this).attr("name"), $(this).attr("isSorted"), isTotal)
})


function sortIconUp() {
    return `<svg id="sortIcon" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-up" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M8 15a.5.5 0 0 0 .5-.5V2.707l3.146 3.147a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L7.5 2.707V14.5a.5.5 0 0 0 .5.5z"/>
            </svg>`
}

function sortIconDown() {
    return `<svg id="sortIcon" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-down" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M8 1a.5.5 0 0 1 .5.5v11.793l3.146-3.147a.5.5 0 0 1 .708.708l-4 4a.5.5 0 0 1-.708 0l-4-4a.5.5 0 0 1 .708-.708L7.5 13.293V1.5A.5.5 0 0 1 8 1z"/>
            </svg>`
}