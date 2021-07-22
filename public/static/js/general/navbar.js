
function getToken() {

	return fetch('/api/auth/gettoken', {
		method: 'GET',
	}).then(function (response) {
		return response.json();
	}).then(function (data) {
		// console.log(data);
		return data.data;

	}).catch(error => {
		console.error(error);
	});
}
window.addEventListener("load", () => {

	// Set navbar for account if logged in
	setTimeout(function () {
		getToken().then(dataResponse => {
			// console.log(dataResponse);
			const myHeaders = new Headers();
			myHeaders.append('Accept', '*/*')
			myHeaders.append('Authorization', 'Bearer ' + dataResponse);

			// console.log(getToken());
			fetch('/api/auth/isloggedin', {
				method: 'POST',
				headers: myHeaders,
			}).then(function (response) {
				return response.json();
			})
				.then(function (data) {
					(function () {
						console.log(data);
						if (data.status === "success") {
							document.getElementById("userArea").innerHTML = 
							`
							<!-- Right elements -->
							<ul class="navbar-nav flex-row">
							<li class="nav-item me-3 me-lg-1">
								<a class="nav-link d-sm-flex align-items-sm-center" href="#">
								<img
									src="https://image.flaticon.com/icons/png/512/4825/4825123.png"
									class="rounded-circle"
									height="22"
									alt=""
									loading="lazy"
								/>
								<strong class="d-none d-sm-block ms-1">`+ data.name_user +`</strong>
								</a>
							</li>
							<li class="nav-item dropdown me-3 me-lg-1">
								<a
								class="nav-link dropdown-toggle hidden-arrow"
								href="#"
								id="navbarDropdownMenuLink"
								role="button"
								data-mdb-toggle="dropdown"
								aria-expanded="false"
								>
								<i class="fas fa-bell fa-lg"></i>
								<span class="badge rounded-pill badge-notification bg-danger">12</span>
								</a>
								<ul
								class="dropdown-menu dropdown-menu-end"
								aria-labelledby="navbarDropdownMenuLink"
								>
								<li>
									<a class="dropdown-item" href="#">Some news</a>
								</li>
								<li>
									<a class="dropdown-item" href="#">Another news</a>
								</li>
								<li>
									<a class="dropdown-item" href="#">Something else here</a>
								</li>
								</ul>
							</li>
							<li class="nav-item dropdown me-3 me-lg-1">
								<a
								class="nav-link dropdown-toggle hidden-arrow"
								href="#"
								id="navbarDropdownMenuLink"
								role="button"
								data-mdb-toggle="dropdown"
								aria-expanded="false"
								>
								<i class="fas fa-chevron-circle-down fa-lg"></i>
								</a>
								<ul
								class="dropdown-menu dropdown-menu-end"
								aria-labelledby="navbarDropdownMenuLink"
								>
									<li>
										<a class="dropdown-item" href="#">My profile</a>
									</li>
										<li>
										<a class="dropdown-item" href="#">Settings</a>
									</li>
										<li>
									<a class="dropdown-item " id="btnLogout" onclick="clickLogout()" href="#">Logout</a>
									</li>
								</ul>
							</li>
							</ul>
							<!-- Right elements -->
							
							`

							document.getElementById("yourArea").innerHTML =
								`<div class="navbar-nav">
									<a class="nav-link active text-light fs-5 me-2" aria-current="page"
										href="#">Your Area</a>
								</div>
								`
							return;
						} else {
							document.getElementById("userArea").innerHTML = 
							`<a href="/account">
								<button class="btn btn-outline-success btn-rounded me-3">Sign in</button>
							</a>`
						}

					})();
				})

		}).catch(function (err) {
				console.log('error: ' + err);
			});
	}, 500);
})

function clickLogout() {
	fetch('/api/auth/logout', {
		method: 'GET',
	}).then(function (response) {
		return response.json();
	})
		.then(function (data) {
			console.log(data);
			(function () {
				if (data.status === "success") {
					window.location.href = "/";
				}

			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});
}
