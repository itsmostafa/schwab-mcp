from colorama import Fore
import requests
from requests.exceptions import HTTPError

from typing import Any
from decimal import Decimal


def get_ohlc_data(ticker, duration):
    endpoint = "https://cloud.iexapis.com/stable/"
    token = "pk_df3dd13bd7aa401e862aefb484dd0058"
    try:
        resp = requests.get(f"{endpoint}/stock/{ticker}/chart/{duration}?token={token}")
        resp.raise_for_status()
    except HTTPError as err:
        print(f"API Error: {err}")

    return resp.json()


def get_stats(data: Any, init_balance: int, stop_loss):
    pnl = list()
    mfe = 0
    mae = 0
    max_drawdown = 0
    max_drawdown_date = ""
    max_profit = 0
    max_profit_date = ""
    cur_balance = init_balance
    start_date = data[0].get("date")
    end_date = data[-1].get("date")
    returns = 0

    for row in data:
        if row.get("open") is not None:
            if stop_loss:
                stop_price = row.get("open") - (row.get("open") * int(stop_loss) / 100)
                if stop_price > row.get("low"):
                    returns = row.get("open") - stop_price
                else:
                    returns = row.get("open") - row.get("close")
            else:
                returns = row.get("open") - row.get("close")
            pnl.append(Decimal(returns).quantize(Decimal(".01")))

            cur_mfe = row.get("high") - row.get("open")
            cur_mae = row.get("open") - row.get("low")

            if sum(pnl) > max_profit:
                max_profit = round(sum(pnl), 2)
                max_profit_date = row["date"]
            if sum(pnl) < max_drawdown:
                max_drawdown = round(sum(pnl), 2)
                max_drawdown_date = row["date"]
            if cur_mfe > mfe:
                mfe = round(cur_mfe, 2)

            if cur_mae > mae:
                mae = round(cur_mae, 2)

            cur_balance += returns

    pos_trades = [trade for trade in pnl if trade > 0]
    neg_trades = [trade for trade in pnl if trade < 0]

    context = {
        "date range": f"{start_date} - {end_date}",
        "pnl": round(sum(pnl), 2),
        "current balance": round(cur_balance, 2),
        "number of trades": len(data),
        "number of positive trades": len(pos_trades),
        "number of negative trades": len(neg_trades),
        "max profit": max_profit,
        "max profit date": max_profit_date,
        "max drawdown": max_drawdown,
        "max drawdown date": max_drawdown_date,
        "mfe": mfe,
        "mae": mae,
    }
    return context


def main():
    header = """
         pppp        aa
         pp  pp    aa  aa 
         pp   pp  aa    aa
         pp ppp   aaaaaaaa
         pp       aa    aa
         pp      aa      aa
         pp      aa      aa

    Welcome to Pure Algo analytics!
    """
    print(Fore.CYAN + header)
    print(Fore.WHITE + "*********************************************")
    init_balance = int(input("What is your initial account balance?"))
    ticker = input("Which stock you like to get stats for?")
    duration = input("What is the date range? [Example: 3m, 6m, 1y, 2y, 5y]")
    stop_loss = input("What is your stop loss? (Percentage)")
    if stop_loss == "":
        stop_loss = None
    ohlc_data = get_ohlc_data(ticker, duration)
    stats = get_stats(ohlc_data, init_balance, stop_loss)

    cur_balance = stats["current balance"]
    pnl_percentage = round(((cur_balance - init_balance) / init_balance) * 100, 2)

    if stats["number of negative trades"] != 0:
        win_loss_ratio = round(
            stats["number of positive trades"] / stats["number of negative trades"], 2
        )
    else:
        win_loss_ratio = 1

    print(Fore.YELLOW + "*********************************************")
    print(f"Here are your results for trades between {stats['date range']}:")
    print("PnL ($):", stats["pnl"], "dollars")
    print("Returns (%):", f"{pnl_percentage}%")
    print("Current Balance:", stats["current balance"])
    print("Total # of trades:", stats["number of trades"])
    print("# of positive trades:", stats["number of positive trades"])
    print("# of negative trades:", stats["number of negative trades"])
    print("Win / Loss Ratio:", win_loss_ratio)
    print("Max Profit:", stats["max profit"])
    print("Max Profit Date:", stats["max profit date"])
    print("Max Drawdown:", stats["max drawdown"])
    print("Max Drawdown Date:", stats["max drawdown date"])
    print("MFE:", stats["mfe"])
    print("MAE:", f'-{stats["mae"]}')


if __name__ == "__main__":
    main()
