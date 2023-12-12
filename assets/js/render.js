

function renderCheck() {
    fetch("/showUser").then(response => response.json()).then(
        function (response) {
            const len = response.user.length
            for (let i = 0; i < len; i++) {

                const userId = response.user[i].ID
                const userName = response.user[i].Name
                renderHtml(userId, userName)
            }


        }
    )
}

function renderHtml(id, name) {

    const tbody = document.querySelector("tbody")
    const row = document.createElement("tr")
    const tdCount = 6
    let tdContent = ""

    for (let i = 0; i < tdCount; i++) {
        tdContent +=
            `
                            <td>
                               <label>
                                    <input class="form-check-input" type="checkbox"  name="punch" value="">
                              </label>                          
                            </td>
            `
    }

    row.innerHTML =
        `
                         <th scope="row">${id}</th>
                         <td>${name}</td>
                         ${tdContent}
       `


    tbody.appendChild(row)

}
