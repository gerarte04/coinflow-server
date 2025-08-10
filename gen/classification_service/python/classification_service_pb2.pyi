from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class GetTextCategoryRequest(_message.Message):
    __slots__ = ("text", "labels")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    text: str
    labels: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, text: _Optional[str] = ..., labels: _Optional[_Iterable[str]] = ...) -> None: ...

class GetTextCategoryResponse(_message.Message):
    __slots__ = ("category",)
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    category: str
    def __init__(self, category: _Optional[str] = ...) -> None: ...
