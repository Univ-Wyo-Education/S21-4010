main:: IO()

main = do putStrLn "Insert the first value: " 
          one <- getLine  
          putStrLn "Insert the second value: "  
          two <- getLine    
          putStrLn "The result is:"
          print ((read one) + (read two))
