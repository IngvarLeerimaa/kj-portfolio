"use strict"
const crosswordSolver = (crossword, words) => {
    if (
        typeof crossword !== "string" ||
        !Array.isArray(words) ||
        words.some((word) => typeof word !== "string")
    ) {
        console.log("Error")
        return "Error"
    }
    // only allow '.', '\n', 0, 1 and 2
    if (!/^[.\n012]+$/.test(crossword)) {
        console.log("Error")
        return "Error"
    }
    // Two-dimensional array with information about words beginning cells in crossword.
    //
    // Each cell contains number of words starting from this cell (0, 1 or 2) or -1 if cell is not available
    const puzzleNumbers = crossword
        .trim()
        .split("\n")
        .map((row) => row.split("").map((cell) => (cell === "." ? -1 : parseInt(cell))))
    const wordsBeginnings = puzzleNumbers
        .map((row, rowIndex) =>
            row.map((cell, colIndex) => ({
                row: rowIndex,
                col: colIndex,
            }))
        )
        .flat()
        .filter((cell) => puzzleNumbers[cell.row][cell.col] > 0)
    if (
        wordsBeginnings.reduce((acc, cell) => acc + puzzleNumbers[cell.row][cell.col], 0) !==
        words.length
    ) {
        console.log("Error")
        return "Error"
    }
    const puzzleWidth = puzzleNumbers[0].length
    if (puzzleNumbers.some((row) => row.length !== puzzleWidth)) {
        console.log("Error")
        return "Error"
    }
    // words repetition
    if (new Set(words).size !== words.length) {
        console.log("Error")
        return "Error"
    }
    // sort words by length (to add the longest to board first)
    words.sort((a, b) => b.length - a.length)
    // Two-dimensional array with information about words placed in crossword.
    //
    // Each cell contains letter if cell is occupied, "" if cell is empty or "." if cell is not available
    const puzzleWords = puzzleNumbers.map((row) => row.map((cell) => (cell === -1 ? "." : "")))
    // Function that checks if it's possible to place word in crossword
    // starting from cell (row, col) in direction (horizontal or vertical)
    const canAddWord = (word, row, col, direction) => {
        var _a
        if (direction === "horizontal" && col + word.length > puzzleNumbers[row].length) {
            return false
        }
        if (direction === "vertical" && row + word.length > puzzleNumbers.length) {
            return false
        }
        for (let i = 0; i < word.length; i++) {
            if (puzzleWords[row][col] !== "") {
                if (puzzleWords[row][col] !== word[i]) {
                    return false
                }
            }
            direction === "horizontal" ? col++ : row++
        }
        // cell after word should be unavailable (or out of the board)
        const afterWordCell = (_a = puzzleNumbers[row]) === null || _a === void 0 ? void 0 : _a[col]
        return afterWordCell === -1 || afterWordCell === undefined
    }
    const addWords = (words) => {
        if (words.length === 0) {
            return true
        }
        for (const word of words) {
            for (const cell of wordsBeginnings) {
                if (puzzleNumbers[cell.row][cell.col] === 0) continue
                if (canAddWord(word, cell.row, cell.col, "horizontal")) {
                    const backupRow = puzzleWords[cell.row].slice()
                    for (let j = 0; j < word.length; j++) {
                        puzzleWords[cell.row][cell.col + j] = word[j]
                    }
                    puzzleNumbers[cell.row][cell.col]--
                    if (addWords(words.filter((w) => w !== word))) {
                        return true
                    }
                    puzzleNumbers[cell.row][cell.col]++
                    puzzleWords[cell.row] = backupRow
                }
                if (canAddWord(word, cell.row, cell.col, "vertical")) {
                    const backupCol = puzzleWords.map((row) => row[cell.col])
                    for (let j = 0; j < word.length; j++) {
                        puzzleWords[cell.row + j][cell.col] = word[j]
                    }
                    puzzleNumbers[cell.row][cell.col]--
                    if (addWords(words.filter((w) => w !== word))) {
                        return true
                    }
                    puzzleNumbers[cell.row][cell.col]++
                    puzzleWords.forEach((row, index) => (row[cell.col] = backupCol[index]))
                }
            }
        }
        return false
    }
    if (!addWords(words)) {
        console.log("Error")
        return "Error"
    }
    const result = puzzleWords.map((row) => row.join("")).join("\n")
    console.log(result)
    return result
}

const crossword = '2001\n0..0\n1000\n0..0'
const words = ['casa', 'alan', 'ciao', 'anta']

crosswordSolver(crossword, words)
