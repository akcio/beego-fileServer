<table class="table table-bordered">
    {{ if .Files}}
        <thead>
            <tr>
                <td>File name</td>
                <td>Upload time</td>
                <td></td>
                <td></td>
            </tr>
        </thead>
        <tbody>
        {{range $key, $val := .Files}}
            <tr>

                <td>{{$val.UserFileName}}</td>
                <td>{{$val.UploadTime}}</td>
                <td><a href="/download/{{$val.UploadTime}}">Download</a></td>
                <td><a href="/delete/{{$val.UploadTime}}">Remove</a></td>
            </tr>
        {{end}}
        </tbody>
    {{end}}

</table>