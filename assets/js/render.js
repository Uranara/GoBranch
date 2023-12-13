function renderCheck() {

    fetch("/showUser").then(response => response.json()).then(
        function (response) {
            const len = response.user.length
            for (let i = 0; i < len; i++) {

                const userId = response.user[i].ID
                const userName = response.user[i].Name

                renderHtml(userId, userName)
            }

            disCheck()
            DataPost()
            showCheck()


        }
    )
}

function renderHtml(id, name) {

    const tbody = document.querySelector("tbody")

    if (!tbody) {
        return
    }
    const row = document.createElement("tr")
    const tdCount = 6
    let tdContent = ""
    const weekList = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat"]


    for (let i = 0; i < tdCount; i++) {
        tdContent +=
            `
                            <td>
                               <label>
                                    <input class="form-check-input ${weekList[i]}" type="checkbox"  name="punch" value="${name}">
                              </label>                          
                            </td>
            `
    }

    let nickname = nameHidden(name)

    row.innerHTML =
        `
                         <th scope="row">${id}</th>
                         <th ><img src="../assets/avatars/${name}.jpg" alt=""></th>
                         <td>${nickname}</td>
                         ${tdContent}
       `


    tbody.appendChild(row)


}

function nameHidden(name) {
    const nameLen = name.length
    let count;
    let asterisk = "";
    let nickname;
    let processName;
    count = nameLen % 3
    if (nameLen > 5) {
        count = nameLen % 5
    }
    for (let i = 0; i < count; i++) {
        asterisk += "*"
    }
    processName = name.slice(0, count) + asterisk
    nickname = processName + name.slice(count + count, nameLen)
    return nickname
}