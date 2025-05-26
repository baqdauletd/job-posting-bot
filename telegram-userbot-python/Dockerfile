# Use official Python image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Copy dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the code
COPY . .

# Set environment variables for Telethon session (optional)
ENV PYTHONUNBUFFERED=1

# Run the bot
CMD ["python", "userbot.py"]
