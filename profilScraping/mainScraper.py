import datetime
import time

import requests
from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
# from profileScraper import scrapData
from webdriver_manager.firefox import GeckoDriverManager

pattern = "https://sharechat.com/profile/"


def isMatch(string, sub_str):
    if string.find(sub_str) == -1:
        return False
    else:
        return True


option = webdriver.FirefoxOptions()
option.headless = True
driver = webdriver.Firefox(executable_path=GeckoDriverManager().install(), options=option)
driver.get("https://sharechat.com")

# Selecting specific language in this gujrati
driver.find_element_by_xpath("//body/div[@id='root']/div[contains(@class,'Pos(f) Start(0) W(100%) T(0) H(100%) Ov(a) Z(100) Bgc($white)')]/div[contains(@class,'')]/main[contains(@class,'Pos(r) Mx(a) Pt($2xl)')]/div[contains(@class,'Py($lg) Px($sm)')]/div/div[4]/div[1]").click()

driver.implicitly_wait(100)
pause_time = 2
last_height = driver.execute_script("return document.body.scrollHeight")
start = datetime.datetime.now()

count = 0
while True:
    count = count + 1
    if count == 5:
        break

    driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
    time.sleep(pause_time)
    new_height = driver.execute_script("return document.body.scrollHeight")
    if new_height == last_height:
        break
    last_height = new_height

# link_tags = driver.find_element_by_css_selector('')
link_tags = driver.find_elements_by_tag_name("a")
print(link_tags)

hrefs = []
for tag in link_tags:
    src = tag.get_attribute('href')
    if isMatch(src, pattern):
        hrefs.append(src)
# print(hrefs)
driver.close()

for link in hrefs:
    print(link)
    # scrapData(link)
    r = requests.post(url="http://localhost:8080/urls/a", json={
        "profile_url": link
    })
    print(r.text)