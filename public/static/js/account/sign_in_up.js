
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

function btnSignUpClick() { 
	let isValidForm = true;
	if(document.getElementById("nameSignUp").checkValidity() == false) {
		isValidForm = false
	}
	if(document.getElementById("emailSignUp").checkValidity() == false) {
		isValidForm = false
	}
	if(document.getElementById("passwordSignUp").checkValidity() == false) {
		isValidForm = false
	}

	if(isValidForm === false) {
		console.log("error invalid")
		return;
	}
	const formData = new FormData();

	formData.append('name', document.getElementById("nameSignUp").value);
    formData.append("email", document.getElementById("emailSignUp").value);
    formData.append("password", document.getElementById("passwordSignUp").value);

	fetch('/api/signup', {
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


function btnSignInClick() { 
	let isValidForm = true;
	if(document.getElementById("emailSignIn").checkValidity() == false) {
		isValidForm = false
	}
	if(document.getElementById("passwordSignIn").checkValidity() == false) {
		isValidForm = false
	}

	if(!isValidForm) {
		console.log("error invalid")
		return;
	}
	const formData = new FormData();

    formData.append("email", document.getElementById("emailSignIn").value);
    formData.append("password", document.getElementById("passwordSignIn").value);

	fetch('/api/auth/signin', {
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
					if(data.message === 'Account not found') {
						document.getElementById("notExistEmail").innerHTML 
							= 	`<p class='text-warning m-0' style='font-size: 12px'>
									This email is not signed up yet, let sign up!
								</p>`;
						return;
					} else if (data.message === "Invalid password") {
						document.getElementById("wrongPass").innerHTML 
							= 	`<p class='text-warning m-0' style='font-size: 12px'>
									password was wrong, try again!
								</p>`;
						return;
					} else {
						console.log(data.message);
					}
				} else if (data.status === 'success') {
					window.location.href = "/"
				}

			})();
		})
}