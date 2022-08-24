import datetime
from zoneinfo import ZoneInfo

now = datetime.datetime.now(ZoneInfo("Asia/Tokyo"))
today = now.strftime("%Y/%m/%d (%a)")

print("Hi!")
print("Today is {}".format(today))
