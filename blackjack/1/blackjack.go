package blackjack
 
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "ace":
        return 11
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4    
    case "five":
        return 5
    case "six":
        return 6   
    case "seven":
        return 7  
    case "eight":
        return 8 
    case "nine":
        return 9 
     case "ten":
        return 10 
     case "jack":
        return 10  
      case "queen":
        return 10  
      case "king":
        return 10         
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	 
     valueC := ParseCard(card1) + ParseCard(card2)
     dealC := ParseCard(dealerCard)
     switch  {
     case valueC > 21:
         return "P"
      case valueC == 21 && dealC >= 10:
         return "S"         
      case valueC == 21:
         return "W"
        case valueC <= 11:
         return "H"    
           case valueC >= 12 && valueC <= 16 && dealC >= 7:
          return "H"    
       case valueC < 21:
         return "S"   
         
     default:
         return "P"
     }
  
}
