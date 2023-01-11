package WeeklyContest

func categorizeBox(length int, width int, height int, mass int) string {
    a := length >= 1e4 || width >= 1e4 || height >= 1e4 || length * width * height >= 1e9
    b := mass >= 100
    
    switch {
    case a && b :
        return "Both"
    case a:
        return "Bulky"
    case b:
        return "Heavy"
    default:
        return "Neither"
    }
}