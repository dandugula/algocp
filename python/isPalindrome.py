def is_palindrome(s):
    cleaned = ''.join(char.lower() for char in s if char.isalnum())

    # Two pointers
    left, right = 0, len(cleaned) - 1
    while left < right:
        if cleaned[left] != cleaned[right]:
            return False
        left += 1
        right -= 1

    return True

