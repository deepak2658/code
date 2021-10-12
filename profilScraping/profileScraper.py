import string
import time

from kafka import KafkaConsumer
from webdriver_manager.firefox import GeckoDriverManager


def isMatch(string, sub_str):
    if (string.find(sub_str) == -1):
        return False
    else:
        return True

pattern = "jpeg"
def scrapData(url):
    import requests
    from selenium import webdriver
    from webdriver_manager.chrome import ChromeDriverManager
    import datetime

    option = webdriver.FirefoxOptions()
    option.headless = True
    driver = webdriver.Firefox(executable_path=GeckoDriverManager().install(), options=option)
#     driver.maximize_window()
    # driver.get("https://in.pinterest.com/")
    print(url)
    driver.get(url)

    driver.implicitly_wait(100)
    pause_time = 2

    profileName = driver.find_element_by_xpath("//div[@class='Whs(nw) Ovx(h) Tov(e) Maw(100%) Fz($fzbutton) Fw($fwbutton) Lh($lhbody) C($darkText)']").text
    profileHandle = driver.find_element_by_xpath("//span[@class='Whs(nw) Ovx(h) Tov(e) Maw(100%)']").text
    profileIcon = driver.find_element_by_xpath("//div[@class='Pos(a) W($8xl) H($8xl) TranslateY(-33%) Bdrs(50%) Bgc($white2) Bgz(cv) Bgr(nr) Bgp(c_t) Bd($bddp)']").value_of_css_property('background-image')
    tagLine = driver.find_element_by_xpath("//div[@class='Py($sm) Mb($xxs) Fw($fwcaption) C($secondaryDark) Fz($fzbutton) Lh($lhmediumCaption) Ta(c)']").text
    followers = None
    following = None

    hrefs = []
    pattern1 = "follower"
    pattern2 = "following"
    followDetails = driver.find_elements_by_tag_name('a')
    for tag in followDetails:

        src = (tag.get_attribute('href'))
        # print(src)
        if isMatch(src, pattern1):  # for followers count
            if followers is None:
               childrens = tag.find_elements_by_xpath(".//*")
               followers = (childrens[0].text)
            hrefs.append(src)
        if isMatch(src, pattern2):  # for following count
            childrens = tag.find_elements_by_xpath(".//*")
            print(childrens[0].text)
            hrefs.append(src)

    # print(profileHandle,profileIcon,profileName,following,followers)

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
    link_tags = driver.find_elements_by_tag_name("img")
    print(link_tags)

    hrefs = []
    for tag in link_tags:
        src = tag.get_attribute('src')
        if isMatch(src, pattern):
            hrefs.append(src)
    driver.close()

    profile ={'profileUrl':url,'profileName':profileName,'profileHandle':profileHandle,'profileIcon':profileIcon,'tagLine':tagLine,'followers':followers}
    print(profile)
    r = requests.post(url="http://localhost:8080/add", json={
        "profile_url": ""+url,
        "profile_name": profileName,
        "profile_handle": profileHandle,
        "profile_icon_url": profileIcon,
        "tag_line": tagLine,
        "followers": followers,
        "post_urls" : hrefs
    })
    print(r.status_code)
    print(r.text)
    return


consumer = KafkaConsumer('prod.trell_crawler_profile',bootstrap_servers=['65.1.9.139:9092'])
for msg in consumer:
    src = (str(msg.value))
    # url = src[2,len(src)-1]
    size = len(src)
    url = (src[2:size-1])
    print(url)
    scrapData(url)

# 1) PYTHON Scpt. ->> go server post request(profile link) ->kafka producer->python script no 2 (consume data)->scape data -> go server (post)
