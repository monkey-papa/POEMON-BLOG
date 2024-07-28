import random
import smtplib
from email.mime.text import MIMEText

from luntan.settings import from_address, wand

def send_email(to_address, content):
    message = MIMEText(content, 'html', 'utf-8')
    message['From'] = from_address
    message['To'] = to_address
    message['subject'] = '评论'
    email = smtplib.SMTP_SSL('smtp.qq.com', 465, 'utf-8')
    email.login(from_address, wand)
    email.sendmail(from_address, to_address, message.as_string())

def send_code(to_address, comment, name):
    content = '''
            <div style="font-family: serif; line-height: 22px; padding: 30px">
      <div style="display: flex; justify-content: center; width: 100%; max-width: 900px; background-image: url('https://Monkey-PaPa.cn/static/assets/Monkey-PaPa11704791811464651.png'); background-size: cover; border-radius: 10px"></div>
      <div style="margin-top: 20px; display: flex; flex-direction: column; align-items: center">
        <div style="margin: 10px auto 20px; text-align: center">
          <div style="line-height: 32px; font-size: 26px; font-weight: bold; color: #000000">嘿！你在 Monkey-PaPa 中收到一条新评论。</div>
          <div style="font-size: 16px; font-weight: bold; color: rgba(0, 0, 0, 0.19); margin-top: 21px">你收到来自 <abc style="color: #c5343e">{0}</abc> 的评论</div>
        </div>
        <div style="min-width: 250px; max-width: 800px; min-height: 128px; background: #f7f7f7; border-radius: 10px; padding: 32px">
          <div>
            <div style="font-size: 18px; font-weight: bold; color: #c5343e">{0}</div>
            <div style="margin-top: 6px; font-size: 16px; color: #000000">
              <p>【Monkey-PaPa】<abc style="color: #c5343e">{1}</abc></p>
            </div>
          </div>

          <a style="width: 150px; height: 38px; background: #ef859d38; border-radius: 32px; display: flex; align-items: center; justify-content: center; text-decoration: none; margin: 40px auto 0" href="http://www.blog.zjh2002.icu" target="_blank" rel="noopener">
            <span style="color: #db214b">查看详情</span>
          </a>
        </div>
        <div style="margin-top: 20px; font-size: 12px; color: black">此邮件由 Monkey-PaPa 自动发出，直接回复无效，有问题请联系站长vx：z15523692545。</div>
      </div>
    </div>
        '''.format(name, comment)
    send_email(to_address, content)
