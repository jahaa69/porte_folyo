from flask import Flask, render_template
import os
import time

app = Flask(__name__)

@app.route('/')
def index():
    seconde=0
    minute=0
    heure=0
    output = ""
    for heure in range (25):
        for minute in range (61):
            for seconde in range (61):
                time.sleep(1)
                os.system("cls")
                output += "heure:{} | minute:{} | seconde:{}<br>".format(heure, minute, seconde)
    return render_template('index.html', output=output)

if __name__ == '__main__':
    app.run()