<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <title>
        {{if .Tittle }}
            {{.Tittle}}s
        {{ else }}
            FileServer
        {{end}}
    </title>
    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/bootstrap-theme.min.css" rel="stylesheet">
    <link href="/static/css/bootstrap-responsive.css" rel="stylesheet">
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <a class="navbar-brand" href="#">FileServer</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
            {{if .IsLogined }}
                <li class="nav-item"><a class="nav-link" href="/upload">Upload file</a></li>
                    <li class="nav-item"><a class="nav-link" href="/">My files</a> </li>
                <li class="nav-item"><a class="nav-link" href="/logout">Logout</a></li>
            {{ else }}
                <li class="nav-item"><a class="nav-link" href="/register">Register</a></li>
                <li class="nav-item"><a class="nav-link" href="/login">Login</a></li>
            {{ end }}
            </ul>
        </div>
    </nav>
    <div class="container">

        {{.LayoutContent}}

    </div>
</body>
<script src="/static/js/jquery.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
</html>