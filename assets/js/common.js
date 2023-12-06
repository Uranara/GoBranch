function DataPost() {
    $("input[name='punch']").click(function () {

        formPOST("#Form", "/form")
        $(this).closest("form").submit()
        $(this).prop("disabled", true)

    })
    formPOST("#addUser", "/add")
}

function formPOST(id, url) {
    $(id).off("submit").submit(function (event) {

        event.preventDefault();
        let formData = $(this).serialize()
        $.ajax({
            url: url,
            method: "post",
            data: formData,
            success: function (response) {
                console.log("ajax:", response)
            },
            error: function (error) {
                console.error("this is ", error)
            }
        })
    })
}

