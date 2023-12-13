function DataPost() {
    $("input[name='punch']").click(function () {

        formPOST("#Form", "/formPost")
        $(this).closest("form").submit()
        $(this).prop("disabled", true)

    })
    // formPOST("#addUser", "/addUser")
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

function delayUpdate(){
    updateDateTime()
}

function updateDateTime() {
    const currentDateTime = new Date();

    const options = {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
        second: 'numeric'
    };
    const weekday = currentDateTime.toLocaleString('zh-CN', {weekday:"long"});
    const formattedDateTime = currentDateTime.toLocaleString('zh-CN', options);

    const datetimeElement = document.getElementById('datetime');
    const weekElement = document.getElementById('weekday');
    if(datetimeElement && weekElement){
        datetimeElement.textContent = formattedDateTime;
        weekElement.textContent = weekday;
    }


}
