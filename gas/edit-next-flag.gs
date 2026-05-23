/**
 * D2:D101 のフラグ列で、値「1」が複数存在しないように制御する。
 * 「1」を新たに入力したら、既存の「1」は自動で「0」に切り替える。
 * トリガー：スプレットシートから > 編集時
 */
function onEdit(e) {
  const TARGET_SHEET = "シート1"; // 対象シート名（環境に合わせて変更）
  const TARGET_COL = 4; // D列
  const START_ROW = 2;
  const END_ROW = 101;
  const FLAG_ON = 1;
  const FLAG_OFF = 0;

  const range = e.range;
  const sheet = range.getSheet();

  if (sheet.getName() !== TARGET_SHEET) return;

  // 編集範囲がE列に被っているかチェック
  const startCol = range.getColumn();
  const endCol = startCol + range.getNumColumns() - 1;
  if (TARGET_COL < startCol || TARGET_COL > endCol) return;

  // 編集範囲が対象行(2〜101)に被っているかチェック
  const editStartRow = range.getRow();
  const editEndRow = editStartRow + range.getNumRows() - 1;
  if (editEndRow < START_ROW || editStartRow > END_ROW) return;

  // 今回の編集で「1」がセットされた行を収集
  const newOnesRows = [];
  for (
    let r = Math.max(editStartRow, START_ROW);
    r <= Math.min(editEndRow, END_ROW);
    r++
  ) {
    const v = sheet.getRange(r, TARGET_COL).getValue();
    if (Number(v) === FLAG_ON) newOnesRows.push(r);
  }

  if (newOnesRows.length === 0) return;

  // 複数セルに同時に「1」が入った場合は最後の1つだけ残す（他は元に戻す/0にする）
  const isSingleCell = range.getNumRows() === 1 && range.getNumColumns() === 1;
  let keepRow;

  if (newOnesRows.length > 1) {
    keepRow = newOnesRows[newOnesRows.length - 1];
    for (const r of newOnesRows) {
      if (r === keepRow) continue;
      sheet.getRange(r, TARGET_COL).setValue(FLAG_OFF);
    }
    SpreadsheetApp.getActive().toast(
      `複数行に「1」が入力されたため、行${keepRow}のみ「1」にしました`,
      "フラグ調整",
      5,
    );
  } else {
    keepRow = newOnesRows[0];
  }

  // 範囲全体を取得し、keepRow 以外で「1」になっている既存行を「0」に
  const colRange = sheet.getRange(
    START_ROW,
    TARGET_COL,
    END_ROW - START_ROW + 1,
    1,
  );
  const values = colRange.getValues();
  let resetCount = 0;

  for (let i = 0; i < values.length; i++) {
    const row = START_ROW + i;
    if (row === keepRow) continue;
    if (Number(values[i][0]) === FLAG_ON) {
      values[i][0] = FLAG_OFF;
      resetCount++;
    }
  }

  if (resetCount > 0) {
    colRange.setValues(values);
    SpreadsheetApp.getActive().toast(
      `行${keepRow}を「1」にし、他の「1」${resetCount}件を「0」に戻しました`,
      "フラグ切替",
      5,
    );
  }
}
