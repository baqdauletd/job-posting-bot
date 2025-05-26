from telethon import TelegramClient, events
from dotenv import load_dotenv
import os

load_dotenv()

API_ID = int(os.getenv("API_ID"))
API_HASH = os.getenv("API_HASH")
SOURCE_CHANNELS = os.getenv("SOURCE_CHANNELS").split(",")
TARGET_CHAT = os.getenv("TARGET_CHAT")
SESSION_NAME = "userbot_session"

client = TelegramClient(SESSION_NAME, API_ID, API_HASH)

@client.on(events.NewMessage(chats=SOURCE_CHANNELS))
async def forward_message(event):
    try:
        await client.send_message(TARGET_CHAT, event.message)
        print(f"Forwarded: {event.message.text[:60]}...")
    except Exception as e:
        print(f"Failed to forward message: {e}")

async def main():
    await client.start()
    print("UserBot is running...")
    await client.run_until_disconnected()

if __name__ == "__main__":
    import asyncio
    asyncio.run(main())
