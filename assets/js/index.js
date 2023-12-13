$(document).ready(
    function () {
        updateDateTime()
        renderCheck()
        setInterval(delayUpdate, 1000);
    }
)


function disCheck() {

    $("input[name ='punch']").prop("disabled", true)
    const day = new Date()
    const week = day.getDay()

    const weekList = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat"]

    const curWeek = weekList[week - 1]
    const preWeek = weekList[week - 2]

    $(`input.${curWeek}`).prop("disabled", false)
    $(`input.${preWeek}`).prop("disabled", false)
}

function showCheck() {
    fetch("/showPunch").then(response => response.json()).then(
        function (response) {

            const punchList = response["punches"]
            const dateCount = {}
            for (let i = 0; i < punchList.length; i++) {
                const userID = response["punches"][i]["UserID"]
                let currentDate = response["punches"][i]["CreatedAt"]

                fetch("/searchUser?id=" + userID).then(response => response.json()).then(function (response) {
                    const username = response["user"]

                    let day = new Date(currentDate)
                    let week = day.getDay()
                    let year = day.getFullYear()
                    let month = day.getMonth() + 1
                    let date = day.getDate()
                    const time = year + "-" + month + "-" + date

                    const timeUser = time + "-" + username

                    if (dateCount[timeUser]) {
                        dateCount[timeUser] += 1
                    } else {
                        dateCount[timeUser] = 1
                    }

                    let inputs = $(`input[value=${username}]`)
                    if (dateCount[timeUser] > 1) {
                        $(inputs[week - 2]).prop("disabled", true).prop("checked", true)
                    }
                    $(inputs[week - 1]).prop("disabled", true).prop("checked", true)
                })
            }

        },
    )
}



