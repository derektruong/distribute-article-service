
const signUpButton = document.getElementById('signUp');
const emailSignUpInput = document.getElementById('emailSignUp');
const signInButton = document.getElementById('signIn');
const btnSignInSubmit = document.getElementById('btnSignInSubmit');
const container = document.getElementById('container');

signUpButton.addEventListener('click', () => {
	container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
	container.classList.remove("right-panel-active");
});

async function btnSignUpClick() { 
	const formData = new FormData();

	formData.append('name', document.getElementById("nameSignUp").value);
    formData.append("email", document.getElementById("emailSignUp").value);
    formData.append("password", document.getElementById("passwordSignUp").value);

	await fetch('/api/signup', {
		method: 'POST',
		body: formData,
	})
		.then(function (response) {
			return response.json();
		})
		.then(function (data) {
			console.log(data);
			(function () {
				if(data.status === 'error') {
					if(data.message === 'Email already exists') {
						document.getElementById("emailExist").innerHTML 
							= 	`<p class='text-warning m-0' style='font-size: 12px'>
									This email is already registered!
								</p>`;
					} else {
						console.log(data.message);
					}
				} else if (data.status === 'success') {
					window.location.href = "/account/welcome"
				}

			})();
		})
}
// document.getElementById("btnSignUpSubmit").addEventListener('click', () => {
	
// })

function helpEmailSignUp() {
	fetch('../statics/data.json')
		.then(function (response) {
			return response.json();
		})
		.then(function (data) {
			(function() {
				let flag = false
				for (var i = 0; i < data.length; i++) {
					if(emailSignUpInput.value === data[i].email) {
						document.getElementById("emailExist").innerHTML 
							= "<p class='text-warning m-0'>This email is already registered!</p>";

						document.getElementById("btnSignUpSubmit").disabled = true;
						flag = true;
					}
					if(flag) break;
					else {
						document.getElementById("emailExist").innerHTML 
							= "";
						document.getElementById("btnSignUpSubmit").disabled = false;
					}
				}
			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});

	
}

btnSignInSubmit.addEventListener("click", () => {
    const formData = new FormData();

    formData.append("email", document.getElementById("emailSignIn").value);
    formData.append(
        "password",
        document.getElementById("passwordSignIn").value
    );
	formData.append('is_login', 'true');

    fetch("/signin", {
        method: "POST",
        body: formData,
    })
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            console.log(data);
            (function () {

				if(data.message === "Email wrong") {
					document.getElementById("notExistEmail").innerHTML 
							= "<p class='text-warning m-0'>"+ data.text +"</p>";
					document.getElementById("wrongPass").innerHTML 
							= "";
					return
				} else if(data.message === "Pass wrong") {
					document.getElementById("wrongPass").innerHTML 
							= "<p class='text-warning m-0'>"+ data.text +"</p>";
					document.getElementById("notExistEmail").innerHTML 
							= "";
					return

				} else if (data.message === "Set cookie successfully") {
                    window.location.href = "/";
                }
            })();
        })
        .catch(function (err) {
            console.log("error: " + err);
        });
});