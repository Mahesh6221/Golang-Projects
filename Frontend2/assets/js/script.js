function logoutbtn() {
    window.location.href = "/logout"
}


document.getElementById("searchQueryInput").addEventListener('input', function () {
    const searchVal = this.value.toLowerCase().trim();
    const searchMessage = document.getElementById("searchMessage");
    const searchBy = document.getElementById("searchBy").value;
    console.log({ searchVal, searchBy });
    searchMessage.style.display = "none"
    userSearch(searchVal, searchBy);




});


function userSearch(searchVal, searchBy) {
    const elements = document.querySelectorAll('.seach_containerofUserPage');
    const searchMessage = document.getElementById("searchMessage");
    const searchMessage2 = document.getElementById("searchMessage2");
    let searchResultFound = false;
    if (searchBy === "select") {
        if (searchVal.trim() == "") {
            console.log("thdfsgfjgdf");
            elements.forEach(function (el) {
                el.style.display = 'flex';
            });
            return;
        }
    }
    if (searchBy === "select") {
        searchMessage.style.display = "flex";
        searchMessage2.style.display = "none";
        elements.forEach(function (el) {
            el.style.display = 'none';
        });
        return
    } else {
        searchMessage.style.display = "none";
        searchMessage2.style.display = "none";
    }




    elements.forEach(function (el) {
        const userName = el.querySelector('.name').textContent.toLowerCase().trim();
        const userEmail = el.querySelector('.email').textContent.toLowerCase().trim();
        const userAge = el.querySelector('.age').textContent.trim();

        let searchResult = false;
        if (searchBy == "name" && userName.startsWith(searchVal)) {
            searchResult = true;
        } else if (searchBy == "email" && userEmail.startsWith(searchVal)) {
            searchResult = true;
        } else if (searchBy == "age" && userAge.includes(searchVal)) {
            searchResult = true;
        }

        if (searchResult) {
            searchResultFound = true;
            el.style.display = 'flex';
        } else {
            el.style.display = 'none';
        }
    });

    if (!searchResultFound) {
        searchMessage2.style.display = "flex";
    } else {
        searchMessage2.style.display = "none";
    }
}



function resetUserSearch() {
    const elements = document.querySelectorAll('.seach_containerofUserPage');
    const searchBy = document.getElementById("searchBy");
    searchBy.value = "select";
    elements.forEach(function (el) {
        el.style.display = 'flex';
    });
}

