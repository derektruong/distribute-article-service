let userArea = `<a href="/signin">
<button class="btn btn-primary me-3">
	Sign in
</button>
</a>`

function getCookie(cname) {
	let name = cname + "=";
	let decodedCookie = decodeURIComponent(document.cookie);
	let ca = decodedCookie.split(';');
	for (let i = 0; i < ca.length; i++) {
		let c = ca[i];
		while (c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
	}
	return "";
}

window.addEventListener('load', () => {

	fetch('/api/authuser', {
		method: 'GET',
	}).then(function (response) {
		return response.json();
	})
		.then(function (data) {
			// console.log(data);
			(function () {
				if (data.message === "granted!") {
					console.log("ok")
					userArea = `
					<ul class="navbar-nav">
						<!-- Avatar -->
						<li class="nav-item dropdown">
							<a class="nav-link dropdown-toggle d-flex align-items-center" href="#"
								id="navbarDropdownMenuLink" role="button" data-mdb-toggle="dropdown"
								aria-expanded="false">
								<img src="https://image.flaticon.com/icons/png/512/4825/4825123.png"
									class="rounded-circle" height="22" alt="" loading="lazy" />
							</a>
							<ul class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
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
					<!-- end avatar -->
					`

					let yourArea = `
					<div class="navbar-nav">
						<a class="nav-link active text-light fs-5 me-2" aria-current="page"
							href="#">Your Area</a>
					</div>
					`
					setNavbar(userArea, yourArea)
				} else {
					setNavbar(userArea, "")
				}

			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});
});

function setNavbar(userArea, yourArea) {
	document.getElementById("navigation").innerHTML =
		`<!-- Navbar -->
	<nav class="shadow-lg navbar navbar-expand-lg navbar-dark fixed-top">
		<!-- Container wrapper -->
		<div class="container">
			<!-- Navbar brand -->
			<a class="navbar-brand me-2" href="" target="_self">
				<img src="../statics/image/logo.png" height="50" alt="" loading="lazy" style="margin-top: -1px" />
			</a>
			<ul class="navbar-nav me-auto mb-2 mb-sm-1">
				<li class="nav-item">
					<a class="nav-link text-light fs-5 fw-bold me-2" href="/">Cypher News</a>
				</li>
			</ul>
			<!-- Toggle button -->
			<button class="navbar-toggler" type="button" data-mdb-toggle="collapse"
				data-mdb-target="#navbarButtonsExample" aria-controls="navbarButtonsExample" aria-expanded="false"
				aria-label="Toggle navigation">
				<i class="fas fa-bars"></i>
			</button>

			<!-- Collapsible wrapper -->
			<div class="collapse navbar-collapse" id="navbarButtonsExample">
				<!-- Left links -->


				<div class="navbar-collapse collapse show" id="navbarNavAltMarkup">
					<div class="navbar-nav">
						<a class="nav-link active text-light fs-5 me-2" aria-current="page"
							href="/headlines">Headlines</a>
					</div>


					<ul class="navbar-nav">
						<!-- Dropdown -->
						<li class="nav-item dropdown">
							<a class="nav-link dropdown-toggle active text-light fs-5" href="#" id="navbarDropdownMenuLink" role="button"
								data-mdb-toggle="dropdown" aria-expanded="false">
								More News
							</a>
							<ul class="dropdown-menu active text-light fs-5" aria-labelledby="navbarDropdownMenuLink">
								<li>
									<a class="dropdown-item" href="/stocks">Stocks</a>
								</li>
								<li>
									<a class="dropdown-item" href="/technology">Technologies</a>
								</li>
								<li>
									<a class="dropdown-item" href="/science">Science</a>
								</li>
								<li>
									<a class="dropdown-item" href="/sport">Sport</a>
								</li>
							</ul>
						</li>
					</ul>

					<div class="navbar-nav">
						<a class="nav-link active text-light fs-5 me-2" aria-current="page"
							href="#">Explore</a>
					</div>

					`
		+
		yourArea
		+
		`

					<ul class="navbar-nav">
						<!-- Dropdown -->
						<li class="nav-item dropdown">
							<a class="nav-link dropdown-toggle active text-light fs-5" href="#" id="navbarDropdownMenuLink" role="button"
								data-mdb-toggle="dropdown" aria-expanded="false">
								About Me
							</a>
							<ul class="dropdown-menu active text-light fs-5" aria-labelledby="navbarDropdownMenuLink">
								<li>
									<a class="dropdown-item" href="http://truongminhduc.herokuapp.com" target="_blank">My portfolio</a>
								</li>
								<li>
									<a class="dropdown-item" href="https://github.com/derektruong" target="_blank">Visit my github</a>
								</li>
							</ul>
						</li>
					</ul>
				</div>

				<!-- Left links -->
				<form class="d-flex input-group w-auto" action="/search" method="GET">
					<input id="search-box" required type="search" class="form-control rounded" placeholder="Search"
						value="{{ .Query }}" name="q" aria-label="Search" aria-describedby="search-addon" />
					<span class="input-group-text text-white border-0" id="search-addon">
						<i class="fas fa-search"></i>
					</span>
				</form>

				<div class="d-flex align-items-center">

					`
		+
		userArea
		+
		`
					
				</div>
			</div>
			<!-- Collapsible wrapper -->
		</div>
		<!-- Container wrapper -->
	</nav>`;

	if (document.getElementById("search-box").value === "{{ .Query }}") {
		document.getElementById("search-box").value = "";
	}
}

function clickLogout() {
	fetch('/api/authlogout', {
		method: 'GET', 
	  }).then(function (response) {
			return response.json();
		})
		.then(function (data) {
			console.log(data);
			(function() {
				if(data.message === "logout successfully") {
					window.location.href = "/";
				}
				
			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});
}
