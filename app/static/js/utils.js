function validateUrl(){
    let urlField = document.getElementById("urlField");
    let url = urlField.value; 

    try {
        url = new URL(url);
    }catch(e){
        this.event.preventDefault();

        urlFieldError = document.getElementById("urlFieldError");
        urlFieldError.innerHTML = "Please insert a valid URL";

        urlField.classList.replace("is-primary", "is-danger");
        
    }
}

function resetError(){
    let urlField = document.getElementById("urlField");

    if(urlField.value == ""){
        urlField.classList.replace("is-danger", "is-primary");
        urlFieldError = document.getElementById("urlFieldError");
        urlFieldError.innerHTML = "";
    }
}

function copyUrlToClipboard() {
    let shortenedUrlField = document.getElementById("shortenedUrlField");
    let copiedUrl = document.getElementById("copiedUrl");

    shortenedUrlField.select();
    shortenedUrlField.setSelectionRange(0, 99999);

    document.execCommand("copy");

    copiedUrl.innerHTML = `Url copied to clipboard<br />`;
    
    setTimeout(() => {
        copiedUrl.innerHTML = "";
    }, 2000);
}