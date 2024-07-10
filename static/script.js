document.addEventListener("DOMContentLoaded",function(){
    const signupForm = document.getElementById("signupForm");
    const signupError = document.getElementById("signupError");

    signupForm.addEventListener("submit",function(event){
        event.preventDefault();

       const formData = new FormData(signupForm);
       const data = {
        username : formData.get("username"),
        password : formData.get("password"),
        email : formData.get("email")
       };
       fetch("/signup",{
        method: "POST",
        headers:{"Content":"application/json"},
        body: JSON.stringify(data)
       }).then(response =>{
        if (response.ok){
            alert("Signup successfull");
            signupForm.reset();
            signupError.style.display = "none";
        }else{
            throw new Error("Signup Failed");
        }
       }).catch(error => {
            signupError.style.display = "block";
            signupError.textContent = error.message;
       });

    });
});