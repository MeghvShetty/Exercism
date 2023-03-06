import requests

r = requests.get("https://xkcd.com/677/")
resp = requests.get("https://imgs.xkcd.com/comics/asshole.png")

# shows you the attribute and objects that can be accessed via the requests
#print(dir(r))

# Help response in the module requests.models Objects:
#print(help(r))

with open('comic.png','wb') as f:
    f.write(r.content)