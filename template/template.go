package template

var Html = `<!DOCTYPE html>
<html>
    <head>
        <title>
            Spotisheet
        </title>
        <script>
            function copy() {
                let code = document.querySelector(".code")
                let btn = document.querySelector(".copy_btn")
                navigator.clipboard.writeText(code.innerHTML).then(() => {
                    btn.innerHTML = "Copied!"
                })
            }
        </script>
        <style>
            .container {
                display: flex;
      justify-content: center;
      font-size: large;
      align-items: center;
      font-family: monospace;
      flex-direction: column;
      margin-top: 2em;
            }
        .code {
        user-select: all;
        max-width: 22rem;
        overflow-x: auto;
        white-space: nowrap;
        background-color: #3c3c3c;
        color: white;
        padding: 0.4rem;
        margin: 2rem;
        }
        .copy_btn {
        margin: 1rem;
        border: none;
        padding: 0.6rem 1rem;
        font-size: large;
        background-color: #1877F2;
        color: white;
        border-radius: 0.6rem;
        cursor: pointer;
        
        }
        </style>

    </head>
  <body>
    <div class="container">
      <div style="margin: 1em;">Copy and paste the code below in the Google sheet</div>
      <div class="code">%s</div>
      <button class="copy_btn" onclick="copy()">Copy</button>
    </div>
</html>
</body>`
