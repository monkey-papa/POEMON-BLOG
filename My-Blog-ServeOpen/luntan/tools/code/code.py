import random
import smtplib
from email.mime.text import MIMEText

from luntan.settings import from_address, wand

def send_email(to_address, content):
    message = MIMEText(content, 'html', 'utf-8')
    message['From'] = from_address
    message['To'] = to_address
    message['subject'] = '验证码'
    email = smtplib.SMTP_SSL('smtp.qq.com', 465, 'utf-8')
    email.login(from_address, wand)
    email.sendmail(from_address, to_address, message.as_string())

def send_code(to_address):
    code = ""
    for i in range(6):
        code += str(random.choice([random.randint(0, 9), chr(random.randint(97, 122)), chr(random.randint(65, 90))]))
    content = '''
            <div style="font-family: serif; line-height: 22px; padding: 30px">
      <div style="display: flex; justify-content: center; width: 100%; max-width: 900px; background-image: url('https://Monkey-PaPa.cn/static/assets/Monkey-PaPa11704791811464651.png'); background-size: cover; border-radius: 10px"></div>
      <div style="margin-top: 20px; display: flex; flex-direction: column; align-items: center">
        <div style="margin: 10px auto 20px; text-align: center">
          <div style="line-height: 32px; font-size: 26px; font-weight: bold; color: #000000">嘿！你在 Monkey-PaPa 中收到一条新消息。</div>
          <div style="font-size: 16px; font-weight: bold; color: rgba(0, 0, 0, 0.19); margin-top: 21px">你收到来自 Monkey-PaPa 的消息</div>
        </div>
        <div style="min-width: 250px; max-width: 800px; min-height: 128px; background: #f7f7f7; border-radius: 10px; padding: 32px">
          <div>
            <div style="font-size: 18px; font-weight: bold; color: #c5343e">Monkey-PaPa</div>
            <div style="margin-top: 6px; font-size: 16px; color: #000000">
              <p>【Monkey-PaPa.cn】<abc style="color: #c5343e">{0}</abc>为本次验证的验证码，请在5分钟内完成验证。为保证账号安全，请勿泄漏此验证码。</p>
            </div>
          </div>

          <a style="width: 150px; height: 38px; background: #ef859d38; border-radius: 32px; display: flex; align-items: center; justify-content: center; text-decoration: none; margin: 40px auto 0" href="https://Monkey-PaPa.cn" target="_blank" rel="noopener">
            <span style="color: #db214b">有朋自远方来</span>
          </a>
        </div>
        <div style="margin-top: 20px; font-size: 12px; color: black">此邮件由 Monkey-PaPa 自动发出，直接回复无效，有问题请联系站长vx：z15523692545。</div>
      </div>
    </div>
        '''.format(code)
    send_email(to_address, content)
    return code
