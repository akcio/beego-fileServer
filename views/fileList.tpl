<a href="/upload">Download file</a>
<table>
    {{range $key, $val := .Files}}
        <tr>

            <td>{{$val.UserFileName}}</td>
            <td>{{$val.UploadTime}}</td>
            <td><a href="/download/{{$val.UploadTime}}">Download</a></td>
            <td><a href="/delete/{{$val.UploadTime}}">Remove</a></td>
        </tr>
    {{end}}
</table>