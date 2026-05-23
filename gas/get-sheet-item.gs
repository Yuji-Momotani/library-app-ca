const SHEET_NAME = "シート1";
const COL_NO = 1;
const COL_TITLE = 3;
const COL_NEXT_FLAG = 4;

function getNextItem() {
  const spreadsheet = SpreadsheetApp.getActiveSpreadsheet();
  const sheet = spreadsheet.getSheetByName(SHEET_NAME);
  if (!sheet) {
    throw new Error("Sheet not found: " + SHEET_NAME);
  }

  const lastRow = sheet.getLastRow();
  if (lastRow < 2) return null;

  const values = sheet.getRange(2, 1, lastRow - 1, COL_NEXT_FLAG).getValues();
  for (let i = 0; i < values.length; i++) {
    const row = values[i];
    if (String(row[COL_NEXT_FLAG - 1]) === "1") {
      const rowNumber = i + 2;
      const sheetUrl =
        spreadsheet.getUrl() +
        "#gid=" +
        sheet.getSheetId() +
        "&range=A" +
        rowNumber;
      return {
        no: row[COL_NO - 1],
        title: row[COL_TITLE - 1],
        row: rowNumber,
        sheetUrl: sheetUrl,
      };
    }
  }

  return null;
}
