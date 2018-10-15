{{if .Error}}
    <div class="alert alert-danger">
        {{.Error}}
    </div>
{{end}}
<form method="post">
    <div class="form-group col-4">
        <label for="login">Username</label>
        <input type="text" class="form-control" id="login" aria-describedby="loginHelp" placeholder="Input login" name="login"/>
    </div>
    <div class="form-group col-4">
        <label for="exampleInputPassword1">Password</label>
        <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password" name="password"/>
    </div>
    <button type="submit" class="btn btn-primary">Login</button>
</form>