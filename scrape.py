import requests

headers = {
    'Metadata-Flavor': 'Google'
}
def scrape(url):
    res = requests.get(url, headers=headers)
    print(url + ': \n' + res.text)
    lines = res.text.split('\n')
    for line in lines:
        if line.endswith('/'): # directory
            scrape(url + line)
        else: # file
            print(url + line + ': \n' + requests.get(url + line, headers=headers).text)
    

    print('\n')

scrape('http://metadata.google.internal/')
