# Building effective Django queries

## Q objects

- Encapsulate keywords
- Accept | (OR), ~ (NOT), & (AND)

```
from django.db.models import Q


Person.objects.filter(
    Q(
    Q(birdate__year__range=(1981, 1996))
    | Q(birthdate__year__range=(1946, 1964))
    & Q(location__postal_code="CA")
)
```

## F objects

- Used to reference fields
- Avoid loading objects into memory
- Good for mathematical executions
- Can be used with annotations
- Can be used with different fields (from the same model)

```
from django.db.models import F

Person.objects.annotate(
    can_donate_on=F("last_donated") + timedelta(days=56)
).filter(can_donate_on__lt=timezone.now())
```

```
from django.db.models import (
    ExpressionWrapper, F, DateTimeField
)

Event.objects.annotate(
    ends_on=ExpressionWrapper(
        F("starts_at") + F("duration"),
        output_field=DateTimeField()
    )
)
```

## Database Functions

- Helps you access database-specific functions directly
- Work for aggregations
- Imnportant for performance

```
from django.db.models import Value
from django.db.models.functions import Concat

Person.objects.annotate(
    full_name=Concat("first_name", Value(" "), "last_name")
).filter(full_name__icontains="nn p").values_list(
    "full_name"
)

# <PersonQuerySet [('Glenn Parker',), (Ann Powers',)]>
```