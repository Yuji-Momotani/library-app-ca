const SLACK_CHANNEL = "#react-100-knocks";
const SLACK_API_URL = "https://slack.com/api/chat.postMessage";

function sendSlackMessage(text) {
  const token =
    PropertiesService.getScriptProperties().getProperty("SLACK_BOT_TOKEN");
  if (!token) {
    throw new Error("SLACK_BOT_TOKEN is not set in Script Properties.");
  }

  const payload = {
    channel: SLACK_CHANNEL,
    text: text,
  };

  const options = {
    method: "post",
    contentType: "application/json; charset=utf-8",
    headers: {
      Authorization: "Bearer " + token,
    },
    payload: JSON.stringify(payload),
    muteHttpExceptions: true,
  };

  const response = UrlFetchApp.fetch(SLACK_API_URL, options);
  const body = JSON.parse(response.getContentText());
  if (!body.ok) {
    throw new Error("Slack API error: " + body.error);
  }
  return body;
}

function notifyDaily() {
  const item = getNextItem();
  if (!item) {
    sendSlackMessage(":warning: 次回フラグ(E列=1)の行が見つかりませんでした。");
    return;
  }

  const today = Utilities.formatDate(new Date(), "Asia/Tokyo", "yyyy-MM-dd");
  const text =
    "<@U06EXKNBT1A> \n" +
    ":books: Golang CleanArchitecture (" +
    today +
    ")\n" +
    "今日の1問: <" +
    item.sheetUrl +
    "|" +
    "*No." +
    item.no +
    " " +
    item.title +
    ">";

  sendSlackMessage(text);
}

function setupTrigger() {
  ScriptApp.getProjectTriggers()
    .filter(function (t) {
      return t.getHandlerFunction() === "notifyDaily";
    })
    .forEach(function (t) {
      ScriptApp.deleteTrigger(t);
    });

  ScriptApp.newTrigger("notifyDaily")
    .timeBased()
    .atHour(9)
    .everyDays(1)
    .inTimezone("Asia/Tokyo")
    .create();
}
