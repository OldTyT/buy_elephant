from aiogram import Bot, types
from aiogram import types
from aiogram.utils import executor
from aiogram.dispatcher import Dispatcher
from aiogram.types import InlineQuery, InputTextMessageContent, InlineQueryResultArticle
from aiogram.contrib.fsm_storage.memory import MemoryStorage

TOKEN = '' #BOT TG Token

bot = Bot(token=TOKEN)
dp = Dispatcher(bot, storage=MemoryStorage())


@dp.message_handler()
async def echo_message(message: types.Message):
    await bot.send_message(message.from_user.id, f'Все говорят "{message.text}", а ты купи слона.',)


@dp.inline_handler()
async def inline_echo(inline_query: InlineQuery):
    if (inline_query.query == ''):
        item = InlineQueryResultArticle(
            id=1,
            description=f'Все говорят ..., а ты купи слона',
            title=f'Введи текст, что бы купить слона',
            thumb_url=f'https://static4.depositphotos.com/1007572/384/i/600/depositphotos_3846873-stock-photo-baby-elephant-running.jpg',
            input_message_content=InputTextMessageContent(f'Введи текст, что бы купить слона')
        )
        await bot.answer_inline_query(inline_query.id, results=[item], cache_time=1)
        return ()
    item = InlineQueryResultArticle(
        id=1,
        description=f'Все говорят "{inline_query.query}", а ты купи слона.',
        title=str(inline_query.query),
        thumb_url=f'https://static4.depositphotos.com/1007572/384/i/600/depositphotos_3846873-stock-photo-baby-elephant-running.jpg',
        input_message_content=InputTextMessageContent(
            f'Все говорят "{inline_query.query}", а ты купи слона.',
            parse_mode=types.ParseMode.HTML)
    )
    await bot.answer_inline_query(inline_query.id, results=[item], cache_time=1)


if __name__ == '__main__':
    executor.start_polling(dp)